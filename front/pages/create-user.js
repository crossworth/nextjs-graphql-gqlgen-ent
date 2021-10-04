import Head from 'next/head';
import Link from 'next/link';
import { useState } from 'react';
import { gql, useMutation } from '@apollo/client';

const CREATE_USER = gql`
    mutation CreateUser($name: String!) {
        createUser(input: {
            name: $name
        }) {
            id
            name
        }
    }
`;

export default function CreateUser() {
  const [name, setName] = useState('');

  const [createUser, { data, loading, error }] = useMutation(CREATE_USER);

  const handleSubmit = (evt) => {
    evt.preventDefault();
    createUser({ variables: { name: name } });
    setName('');
  };

  if (loading) return 'Submitting...';
  if (error) return `Submission error! ${error.message}`;

  return (
    <div>
      <Head>
        <title>Create user</title>
      </Head>
      <div>
        {data && <div>
          User created <br/>
          {data.createUser.id} - {data.createUser.name}
          <br/>
          <Link href={`/users/${data.createUser.id}`}>{`/users/${data.createUser.id}`}</Link>
          <br/>
          <br/>
        </div>}

        <form onSubmit={handleSubmit}>
          <input type="text" name="name" placeholder="User name"
                 value={name}
                 onChange={e => setName(e.target.value)}
          />
          <button type="submit">Save</button>
        </form>
      </div>
    </div>
  );
}