import type { AppProps } from "next/app";
import { SWRConfig } from "swr";
import { fetcher } from "../utils";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <SWRConfig
      value={{
        fetcher: fetcher,
      }}
    >
      <Component {...pageProps} />
    </SWRConfig>
  );
}
export default MyApp;
