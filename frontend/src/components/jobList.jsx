import React, { useState, useEffect } from "react";
import { VStack, Text, HStack, Badge } from "@chakra-ui/react";

const JobDisplay = () => {
    const [jobs, setJobs] = useState([]);

    useEffect(() => {
        fetchJobs();
        const ws = new WebSocket("ws://localhost:8080/ws");

        ws.onmessage = (event) => {
            const job = JSON.parse(event.data);
            setJobs((prevJobs) => {
                const index = prevJobs.findIndex((j) => j.id === job.id);
                if (index !== -1) {
                    // Update existing job
                    const updatedJobs = [...prevJobs];
                    updatedJobs[index] = job;
                    return updatedJobs;
                } else {
                    // Add new job
                    return [...prevJobs, job];
                }
            });
        };

        ws.onclose = () => {
            console.log("WebSocket connection closed");
        };

        return () => {
            ws.close();
        };
    }, []);

    const fetchJobs = async () => {
        const response = await fetch("http://localhost:8080/jobs");
        const data = await response.json();
        setJobs(data);
    };

    return (
      <VStack>
      {jobs.map((job) => (
          <HStack key={job.id} w="full" p={4} borderWidth={1} borderRadius="md">
              <Text>{job.name}</Text>
              <Text>{job.duration} Seconds</Text>
              <Badge colorScheme={job.status === "completed" ? "green" : job.status === "running" ? "yellow" : "red"}>
                  {job.status}
              </Badge>
          </HStack>
      ))}
  </VStack>
    );
};

export default JobDisplay;
