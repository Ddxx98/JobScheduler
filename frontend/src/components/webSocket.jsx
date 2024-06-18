// src/WebSocketComponent.jsx
import React, { useEffect } from 'react';
import { useSnackbar } from 'notistack';

const WebSocketComponent = () => {
  const { enqueueSnackbar } = useSnackbar();

  useEffect(() => {
    const socket = new WebSocket('http://localhost:8080/ws');

    socket.onmessage = (event) => {
      const job = JSON.parse(event.data);
      enqueueSnackbar(`Job ${job.name} is now ${job.status}`, { variant: 'info' });
    };

    return () => socket.close();
  }, [enqueueSnackbar]);

  return null;
};

export default WebSocketComponent;
