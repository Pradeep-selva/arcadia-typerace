const URL = "localhost:5500";
export const BASE_URL = `http://${URL}`;
const SOCKET_URL = `ws://${URL}`;

export const DEPLOYED_URL = window.location.origin;

export const VAL_TYPES = {
  join: "join",
  create: "create"
};

export const API_ENDPOINTS = {
  validate: (valType: keyof typeof VAL_TYPES, roomId: string) =>
    `/validate/${valType}/${roomId}`,
  socket: (roomId: string) => `${SOCKET_URL}/api/ws/${roomId}`
};
