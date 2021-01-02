import { DEPLOYED_URL } from "@/configs";

export const generateUID = (): string => {
  let firstPart: number | string = (Math.random() * 46656) | 0;
  let secondPart: number | string = (Math.random() * 46656) | 0;
  firstPart = ("000" + firstPart.toString(36)).slice(-3);
  secondPart = ("000" + secondPart.toString(36)).slice(-3);
  return firstPart + secondPart;
};

export const copyToClipBoard = async (content: string, callback: Function) => {
  try {
    await navigator.clipboard.writeText(content);
    callback();
  } catch (error) {
    console.log(error);
  }
};

export const goHome = () => window.location.replace(DEPLOYED_URL);
