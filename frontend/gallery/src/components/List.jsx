import { usePosts } from "../context/PostsContext";
import { motion } from "motion/react";
import Masonry from "react-masonry-css";
import styles from "./List.module.css";
import Post from "./Post";
import Spinner from "./Spinner";

function List() {
	const { posts, isLoading } = usePosts();

	const breakpointColumnsObj = {
		default: 4,
		1100: 3,
		700: 2,
		500: 1,
	};

	if (!posts || posts.length === 0) {
		return (
			<div className="noPostsWrapper">
				<p className="noPosts">There are no posts 🥲</p>
			</div>
		);
	}

	return isLoading ? (
		<Spinner />
	) : (
		<Masonry
			breakpointCols={breakpointColumnsObj}
			className={styles.myMasonryGrid}
			columnClassName={styles.myMasonryGridColumn}
		>
			{[...posts].reverse().map((post) => (
				<motion.div
					key={post.id}
					initial={{ opacity: 0, y: 50 }}
					animate={{ opacity: 1, y: 0 }}
					transition={{
						duration: 0.5,
					}}
				>
					<Post post={post} />
				</motion.div>
			))}
		</Masonry>
	);
}

export default List;
