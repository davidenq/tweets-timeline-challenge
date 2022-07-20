import styles from '../../styles/card.module.css'
export default function Card({ user, tweets }) {
  return (
    <div className={styles.containerCard}>
      <div className={styles.containerPhotoTimeline}>
        <div className={styles.containerImg}>
          <img src={user.profile_image_url} />
        </div>
        <div className={styles.containerDescription}>
          <h6>{user.name} [@{user.screen_name}]</h6>
          <h4>Description</h4>
          <p>{user.description}</p>
        </div>
      </div>

      <div className={styles.containerNameDescription}>
      </div>

      <div className={styles.containerTimeline}>
        {tweets.data.map(tweet => {
          return (
            <div className={styles.containerTweet} id={tweet.id}>
              <p>{tweet.text}</p>
              {tweet.extended_entities.media != undefined ? tweet.extended_entities.media.map(m => (
                <img src={m.media_url_https} width="50px" />
              )) : <></>
              }
            </div>
          )
        })}
      </div>
    </div>
  )
}