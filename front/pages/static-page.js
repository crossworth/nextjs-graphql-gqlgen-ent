import Head from 'next/head';

export default function StaticPage() {
  return (
    <div>
      <Head>
        <title>Static page</title>
      </Head>
      <div>
        This page is static!
      </div>
    </div>
  );
}