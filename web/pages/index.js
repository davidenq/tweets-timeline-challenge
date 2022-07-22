import Head from 'next/head'
import Card from './components/Card'
import styles from '../styles/Home.module.css'
import { useState, useEffect } from 'react'
import getConfig from 'next/config'
const { publicRuntimeConfig } = getConfig()

async function timelineFetch(username) {
  console.log(publicRuntimeConfig)
  const res = await fetch(`${publicRuntimeConfig.apiEndpoint}:${publicRuntimeConfig.apiPort}/api/v1/users/${username}/timeline`)
  const tweets = await res.json()
  const user = tweets.data[0].user
  return {
    tweets: tweets,
    user: user
  }
}

export default function Home() {
  const [username, setUsername] = useState('')
  const [tweetsTimeLine, setTweetsTimeLine] = useState()
  const [user, setUser] = useState(undefined)
  const [searchAction, setSearchAction] = useState(false)
  useEffect(() => {
    if (searchAction) {
      timelineFetch(username)
        .then(({ user, tweets }) => {
          setUser(user)
          setTweetsTimeLine(tweets)
          setSearchAction(false)
        })
    }
  });

  return (
    <div className={styles.container}>
      <Head>
        <title>Tweets timeline by user</title>
        <meta name="description" content="tweets timeline challenge" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <h1>Tweets timeline by user</h1>
      <input
        value={username}
        placeholder='search by twitter username'
        onChange={e => setUsername(e.target.value)} />
      <button disabled={searchAction} onClick={setSearchAction}>Search</button>
      {user === undefined ? "search by username" : <Card user={user} tweets={tweetsTimeLine} />}
    </div>
  )
}