import { SECRETS } from "@/secrets";

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
  socket: (roomId: string, userName: string) =>
    `${SOCKET_URL}/api/ws/x-api-key@${SECRETS.X_API_KEY}/${roomId}/${userName}`
};
