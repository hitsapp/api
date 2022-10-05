export const IS_PROD = process.env.NODE_ENV === "production";
export const BASE_URI = IS_PROD ? "https://hits.hop.sh" : "http://localhost";
export const API_URL = `${BASE_URI}/v1`;
export const WEBSITE_REGEX =
  /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)/;
