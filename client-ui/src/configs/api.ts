import { SECRETS } from "@/secrets";
import { PRODUCTION } from "./environment";

const URL = PRODUCTION ? SECRETS.API_URL : "localhost:5500";
export const BASE_URL = `${PRODUCTION ? "https" : "http"}://${URL}`;
const SOCKET_URL = `${PRODUCTION ? "wss" : "ws"}://${URL}`;

export const DEPLOYED_URL = PRODUCTION
  ? "https://arcadia-typerace.web.app"
  : "http://localhost:8080/";

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
