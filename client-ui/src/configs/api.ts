export const BASE_URL = "http://localhost:5500";

export const DEPLOYED_URL = window.location.origin;

export const VAL_TYPES = {
  join: "join",
  create: "create"
};

export const API_ENDPOINTS = {
  validate: (valType: keyof typeof VAL_TYPES, roomId: string) =>
    `/validate/${valType}/${roomId}`,
  socket: (roomId: string) => `/ws/${roomId}`
};
