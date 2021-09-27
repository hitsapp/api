export const IS_PROD = process.env.NODE_ENV === "production";
export const BASE_URI = IS_PROD ? "https://api.hits.link" : "http://localhost";
export const API_URL = `${BASE_URI}/v1`;
