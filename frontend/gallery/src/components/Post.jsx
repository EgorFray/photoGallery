import { Link } from "react-router-dom";
import styles from "./Post.module.css";

function Post({ post }) {
	const { id, image, description } = post;
	return (
		<li key={id} className={styles.post}>
			<Link to={`posts/${id}`}>
				<img
					className={styles.postImg}
					src={`http://localhost:8080${image}`}
					alt="Something beeautiful here"
				/>
				<div className={styles.postBox}>
					<p className={styles.postDesc}>{description} </p>
				</div>
			</Link>
		</li>
	);
}

export default Post;
