/* eslint-disable react/display-name */
import Document, {
  Head,
  Html,
  NextScript,
  Main,
  DocumentContext,
} from "next/document";
import { ServerStyleSheet } from "styled-components";

const Page = () => {
  return (
    <Html lang='en'>
      <Head>
        <meta charSet='UTF-8' />
        <title>Hits</title>
        <meta name='description' content='Count your hits!' />
      </Head>
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  );
};

export default class extends Document {
  static async getInitialProps(ctx: DocumentContext): Promise<{
    styles: JSX.Element;
    html: string;
    head?: (JSX.Element | null)[] | undefined;
  }> {
    const sheet = new ServerStyleSheet();
    const originalRenderPage = ctx.renderPage;

    try {
      ctx.renderPage = () =>
        originalRenderPage({
          enhanceApp: (App) => (props) =>
            sheet.collectStyles(<App {...props} />),
        });

      const initialProps = await Document.getInitialProps(ctx);
      return {
        ...initialProps,
        styles: (
          <>
            {initialProps.styles}
            {sheet.getStyleElement()}
          </>
        ),
      };
    } finally {
      sheet.seal();
    }
  }

  render(): JSX.Element {
    return <Page />;
  }
}
