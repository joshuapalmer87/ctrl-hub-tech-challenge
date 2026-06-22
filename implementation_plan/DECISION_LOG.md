1. Using a single test instance for testing purposes
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