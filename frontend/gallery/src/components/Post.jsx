import styles from "./Post.module.css";

function Post({ post }) {
	const { id, image, description } = post;
	return (
		<li key={id} className={styles.post}>
			<img
				className={styles.postImg}
				src={`http://localhost:8080${image}`}
				alt="Something beeautiful here"
			/>
			<div className={styles.postBox}>
				<p className={styles.postDesc}>{description} </p>
			</div>
		</li>
	);
}

export default Post;
