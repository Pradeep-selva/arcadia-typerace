export const Events = {
  JOIN: "join",
  WON: "won",
  TYPING: "typing"
};

export interface EventResponse {
  event: string;
  data: string;
  userName: string;
}
