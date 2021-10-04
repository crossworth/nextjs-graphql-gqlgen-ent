import Head from 'next/head';
import Link from 'next/link';

export default function Home() {
  return (
    <div>
      <Head>
        <title>Home</title>
      </Head>
      <div>
        Home
        <br/>
        <Link href="/static-page">
          <a>Static page</a>
        </Link>
        <br/>
        <Link href="/create-user">
          <a>Create user</a>
        </Link>
      </div>
    </div>
  );
}
