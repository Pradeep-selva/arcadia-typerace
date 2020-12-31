import { API_ENDPOINTS, VAL_TYPES } from "@/configs";
import Axios from "./axios";

interface Response {
  data: string;
  ok: boolean;
}

export const validateOperation = (
  valType: keyof typeof VAL_TYPES,
  roomId: string
) =>
  Axios.getInstance()
    .get<Response>(API_ENDPOINTS.validate(valType, roomId))
    .then((response) => response.data);
