import { VAL_TYPES } from "./api";

export const ROUTES = {
  home: "/",
  validate: "/validate",
  room: (roomId: string = ":id") => `/rooms/${roomId}`
};
