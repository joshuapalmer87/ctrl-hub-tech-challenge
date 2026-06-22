package httpserver_test

import (
	"context"
	"ctrl-hub-technical-challenge/pkg/core/equipment"
	"ctrl-hub-technical-challenge/pkg/core/exposure"
	"ctrl-hub-technical-challenge/pkg/core/model"
	"ctrl-hub-technical-challenge/pkg/core/user"
	"ctrl-hub-technical-challenge/pkg/httpserver"
	"ctrl-hub-technical-challenge/pkg/storage"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const serverAddr = "http://localhost:8090" // TODO - construct from config when added

// TODO - this should be superceded by direct access to the storage medium
var exposureStore *storage.Service

func TestMain(m *testing.M) {
	userService := user.NewService()
	equipmentService := equipment.NewService()
	exposureStorage := storage.NewService()
	exposureStore = exposureStorage
	exposureService := exposure.NewService(userService, equipmentService, exposureStorage)
	server := httpserver.NewHttpServer(exposureService)

	// Serve blocks (http.ListenAndServe), so run it in the background.
	go func() {
		err := server.Serve()
		if err != nil {
			panic(err)
		}
	}()

	// The server boots asynchronously, so retry until it accepts
	// connections, giving up after a short timeout.
	var resp *http.Response
	var err error
	url := serverAddr + "/ping"
	for attempt := 0; attempt < 50; attempt++ {
		resp, err = http.Get(url)
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	defer resp.Body.Close()

	// Run tests
	code := m.Run()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		os.Stderr.WriteString("server shutdown: " + err.Error() + "\n")
	}

	os.Exit(code)
}

func TestServerPing(t *testing.T) {
	url := serverAddr + "/ping"

	resp, err := http.Get(url)
	require.NoError(t, err, "could not reach %s", url)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPostExposure(t *testing.T) {
	cleanStorage()

	url := serverAddr + "/exposure"

	body := `{
		"equipment_id": "2e85d43d-dd9b-4e8d-b2ce-97b8d7d69d49",
  		"duration": 480,
		"user_id": "713be58e-0d79-4df2-a85c-9f44ca513a7d"
	}`

	// Test response
	startTime := time.Now()
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	endTime := time.Now()
	require.NoError(t, err, "could not reach %s", url)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	var exposureResult model.Exposure
	err = json.NewDecoder(resp.Body).Decode(&exposureResult)
	require.NoError(t, err, "could not parse exposure")

	assert.Equal(t, "AirCat - Drill - 4337", exposureResult.Equipment.Name)
	assert.Equal(t, 2.1, exposureResult.Equipment.VibrationMagnitude)
	assert.Equal(t, "Bobby Tables", exposureResult.User.Name)
	assert.Equal(t, 480, exposureResult.DurationMinutes)
	assert.Equal(t, 2.1, exposureResult.A8)
	assert.Equal(t, 71.0, exposureResult.Points)

	// Test stored values
	storedExposure, ok := exposureStore.ExposureMap[exposureResult.ID]
	require.True(t, ok)
	assert.Equal(t, exposureResult, storedExposure)

	storedExposuresWithTime, ok := exposureStore.UserExposureMap["713be58e-0d79-4df2-a85c-9f44ca513a7d"]
	require.True(t, ok)
	assert.Len(t, storedExposuresWithTime, 1)
	assert.Equal(t, exposureResult, storedExposuresWithTime[0].Exposure)
	assert.WithinRange(t, storedExposuresWithTime[0].ExposureTime, startTime, endTime)

}

func TestGetAllExposure(t *testing.T) {
	cleanStorage()
	url := serverAddr + "/exposure"

	// TODO - add some fixtures here for creating data and validating it
	exp1 := model.Exposure{
		ID: "3e85d43d-dd9b-4e8d-b2ce-97b8d7d69d49",
		Equipment: model.EquipmentItem{
			ID:                 "2e85d43d-dd9b-4e8d-b2ce-97b8d7d69d49",
			Name:               "AirCat - Drill - 4337",
			VibrationMagnitude: 2.1,
		},
		User: model.User{
			ID:   "1e85d43d-dd9b-4e8d-b2ce-97b8d7d69d49",
			Name: "Bobby Tables",
		},
		DurationMinutes: 5,
		A8:              3.5,
		Points:          5.6,
	}
	exposureStore.ExposureMap[exp1.ID] = exp1
	exp2 := model.Exposure{
		ID: "3e85d43d-dd9b-4e8d-b2ce-97b8d7d69d48",
		Equipment: model.EquipmentItem{
			ID:                 "2e85d43d-dd9b-4e8d-b2ce-97b8d7d69d48",
			Name:               "AirCat - Drill - 4338",
			VibrationMagnitude: 2.8,
		},
		User: model.User{
			ID:   "1e85d43d-dd9b-4e8d-b2ce-97b8d7d69d48",
			Name: "Bobby Tables",
		},
		DurationMinutes: 8,
		A8:              3.8,
		Points:          5.8,
	}
	exposureStore.ExposureMap[exp2.ID] = exp2

	resp, err := http.Get(url)
	require.NoError(t, err, "could not reach %s", url)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var exposures []model.Exposure
	err = json.NewDecoder(resp.Body).Decode(&exposures)
	require.NoError(t, err, "could not parse exposures")
	require.Len(t, exposures, 2)
	assert.Equal(t, exp1, exposures[0])
	assert.Equal(t, exp2, exposures[1])
}

func TestGetExposure(t *testing.T) {
	cleanStorage()
	expID1 := "3e85d43d-dd9b-4e8d-b2ce-97b8d7d69d49"
	url := serverAddr + "/exposure/" + expID1

	exp1 := model.Exposure{
		ID: expID1,
		Equipment: model.EquipmentItem{
			ID:                 "2e85d43d-dd9b-4e8d-b2ce-97b8d7d69d49",
			Name:               "AirCat - Drill - 4337",
			VibrationMagnitude: 2.1,
		},
		User: model.User{
			ID:   "1e85d43d-dd9b-4e8d-b2ce-97b8d7d69d49",
			Name: "Bobby Tables",
		},
		DurationMinutes: 5,
		A8:              3.5,
		Points:          5.6,
	}
	exposureStore.ExposureMap[exp1.ID] = exp1
	exp2 := model.Exposure{
		ID: "3e85d43d-dd9b-4e8d-b2ce-97b8d7d69d48",
		Equipment: model.EquipmentItem{
			ID:                 "2e85d43d-dd9b-4e8d-b2ce-97b8d7d69d48",
			Name:               "AirCat - Drill - 4338",
			VibrationMagnitude: 2.8,
		},
		User: model.User{
			ID:   "1e85d43d-dd9b-4e8d-b2ce-97b8d7d69d48",
			Name: "Bobby Tables",
		},
		DurationMinutes: 8,
		A8:              3.8,
		Points:          5.8,
	}
	exposureStore.ExposureMap[exp2.ID] = exp2

	resp, err := http.Get(url)
	require.NoError(t, err, "could not reach %s", url)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var exposureResult model.Exposure
	err = json.NewDecoder(resp.Body).Decode(&exposureResult)
	require.NoError(t, err, "could not parse exposure")
	assert.Equal(t, exp1, exposureResult)
}

func cleanStorage() {
	exposureStore.ExposureMap = make(map[string]model.Exposure)
}
