import { usePosts } from "../context/PostsContext";
import Masonry from "react-masonry-css";
import styles from "./List.module.css";

function List() {
	const { posts } = usePosts();

	const breakpointColumnsObj = {
		default: 4,
		1100: 3,
		700: 2,
		500: 1,
	};

	return posts ? (
		<Masonry
			breakpointCols={breakpointColumnsObj}
			className={styles.myMasonryGrid}
			columnClassName={styles.myMasonryGridColumn}
		>
			{[...posts].reverse().map((post) => (
				<li key={post.ID} className={styles.post}>
					<img
						className={styles.postImg}
						src={`http://localhost:8080${post.image}`}
						alt="Something beeautiful here"
					/>
					{/* HERE WE'LL ADD DATE IN THE FUTURE */}
					<div className={styles.postBox}>
						<p className={styles.postDesc}>{post.description} </p>
					</div>
				</li>
			))}
		</Masonry>
	) : (
		<div className="noPostsWrapper">
			<p className="noPosts">There are no posts 🥲</p>
		</div>
	);
}

export default List;
