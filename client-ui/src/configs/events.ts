export const Events = {
  JOIN: "join",
  WON: "won",
  TYPING: "typing",
  RECEIVED_TEST: "receivedTest"
};

export interface EventResponse {
  event: string;
  data: string;
  userName: string;
}
