import type { AppProps } from "next/app";
import { SWRConfig } from "swr";
import { fetcher, darkTheme } from "../utils";
import { ThemeProvider, createGlobalStyle } from "styled-components";

const GlobalStyle = createGlobalStyle`
  body {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    background: ${({ theme }) => theme.layoutDarkest};
    color: ${({ theme }) => theme.textLightest};
    font-family: 'Manrope', sans-serif;
    font-weight: 500;
  }
`;

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ThemeProvider theme={darkTheme}>
      <GlobalStyle />
      <SWRConfig
        value={{
          fetcher: fetcher,
        }}
      >
        <Component {...pageProps} />
      </SWRConfig>
    </ThemeProvider>
  );
}
export default MyApp;
