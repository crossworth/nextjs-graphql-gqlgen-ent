import { ApolloProvider } from '@apollo/client';
import { client } from '../service/api';

function App({ Component, pageProps }) {
  return (
    <ApolloProvider client={client}>
      <Component {...pageProps} />
    </ApolloProvider>
  );
}

export default App;
