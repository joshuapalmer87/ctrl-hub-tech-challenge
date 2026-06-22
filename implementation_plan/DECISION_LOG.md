1. Using a single test server for testing purposes
   * PROS 
     * Removes overhead of setup and teardown for each individual test
     * Makes it easier to use in-memory data source
   * CONS
     * Need to be aware of data cross-contamination in tests

2. Using in-memory data source for initial implementation
    * PROS
      * Speed
    * CONS
      * No persistence if server shuts down (only viable as tech challenge!)

3. On implementing the summary endpoint, it was noted no time data was supplied. Given the spec is fixed
we should assume that the time of posting is the end of the time occurred. A better fix would be to alter
the endpoint to take in this data to ensure compliance. Given the limitations, we will return all relevant
exposures that finish within the time frame instead.
    * CONS
      * Likely to provide incomplete data over certain time frames
    * PROS
      * Sticks with the provided spec
      * Unlikely to a problem with 24 hour windows