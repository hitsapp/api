import type { AppProps } from "next/app";
import { SWRConfig } from "swr";
import { fetcher, darkTheme } from "../utils";
import { ThemeProvider } from "styled-components";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ThemeProvider theme={darkTheme}>
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
