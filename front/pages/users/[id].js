import Head from 'next/head';
import { gql } from '@apollo/client';
import { client } from '../../service/api';

const GET_USER = gql`
    query GetUser($id: ID!) {
        user(id: $id) {
            id
            name
        }
    }
`;

export async function getServerSideProps(context) {
  const { id } = context.query;

  try {
    const { data } = await client.query({
      query: GET_USER,
      variables: {
        id
      }
    });

    return {
      props: {
        user: data.user,
      },
    };
  } catch (error) {
    return {
      props: {
        user: null,
        error: error.toString()
      },
    };
  }
}

export default function User({ user, error }) {
  if (error) {
    return (
      <div>
        <Head>
          <title>Error</title>
        </Head>
        <div>
          {error}
        </div>
      </div>
    );
  }

  return (
    <div>
      <Head>
        <title>User - {user.name}</title>
      </Head>
      <div>
        User with id {user.id} and name {user.name}
      </div>
    </div>
  );
}