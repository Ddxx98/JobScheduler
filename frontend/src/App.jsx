import React from "react";
import { ChakraProvider, Box, Heading } from "@chakra-ui/react";
import JobDisplay from "./components/jobList";
import JobSubmission from "./components/jobForm";

function App() {
    return (
        <ChakraProvider>
            <Box p={5}>
                <Heading mb={6}>Job Scheduler</Heading>
                <JobSubmission />
                <JobDisplay />
            </Box>
        </ChakraProvider>
    );
}

export default App;
