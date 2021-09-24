import { DefaultTheme } from "styled-components";
declare module "styled-components" {
  export interface DefaultTheme {
    layoutDarkest: string;
    layoutDark: string;
    layoutLittleLessDark: string;

    textLightest: string;
    textLight: string;
    textDarker: string;
    textDarkest: string;

    success: string;
    error: string;
    info: string;
  }
}

export const darkTheme: DefaultTheme = {
  layoutDarkest: "#202225",
  layoutDark: "#2F3136",
  layoutLittleLessDark: "#36383E",

  textLightest: "#FFFFFF",
  textLight: "#dedede",
  textDarker: "#acacac",
  textDarkest: "#9E9E9E",

  success: "#2fd671",
  error: "#ff4f4f",
  info: "#f1fa8c",
};
