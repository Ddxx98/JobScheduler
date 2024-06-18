import React, { useState } from "react";
import { VStack, Input, Button } from "@chakra-ui/react";

const JobSubmission = () => {
    const [name, setName] = useState("");
    const [duration, setDuration] = useState("");

    const submitJob = async () => {
        const response = await fetch("http://localhost:8080/jobs", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ name, duration: parseInt(duration) }),
        });
        if (response.ok) {
            setName("");
            setDuration("");
        }
    };

    return (
        <VStack>
            <Input
                placeholder="Job Name"
                value={name}
                onChange={(e) => setName(e.target.value)}
            />
            <Input
                placeholder="Duration (ms)"
                type="number"
                value={duration}
                onChange={(e) => setDuration(e.target.value)}
            />
            <Button onClick={submitJob}>Submit Job</Button>
        </VStack>
    );
};

export default JobSubmission;
