export const ROUTES = {
  home: "/",
  validate: "/validate",
  room: (roomId = ":id") => `/rooms/${roomId}`
};

export const ROUTE_NAMES = {
  home: "Home",
  validate: "Validate",
  room: "Room"
};
