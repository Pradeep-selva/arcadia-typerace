export const Events = {
  JOIN: "join",
  WON: "won",
  TYPING: "typing",
  RECEIVED_TEST: "receivedTest",
  INITIATE_START: "initiateStart"
};

export interface EventResponse {
  event: string;
  data: string;
  userName: string;
}
