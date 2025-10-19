import { Link } from "react-router-dom";
import styles from "./Post.module.css";

function Post({ post }) {
	const { id, image } = post;
	return (
		<li key={id} className={styles.post}>
			<Link to={`/app/posts/${id}`}>
				<img
					className={styles.postImg}
					src={`${import.meta.env.VITE_BACKEND_URL}${image}`}
					alt="Something beeautiful here"
				/>
				{/* <div className={styles.postBox}>
					<p className={styles.postDesc}>{description} </p>
				</div> */}
			</Link>
		</li>
	);
}

export default Post;
