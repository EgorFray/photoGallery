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

	return (
		<>
			{isLoading ? (
				<Spinner />
			) : !Array.isArray(posts) || posts.length === 0 ? (
				<div className="noPostsWrapper">
					<p className="noPosts">You have no posts yet</p>
					<motion.p
						className="noPostsShow"
						initial={{ clipPath: "inset(0 100% 0 0)" }}
						animate={{ clipPath: "inset(0 0% 0 0)" }}
						transition={{ duration: 1.2, ease: "easeInOut", delay: 1.5 }}
					>
						Maybe it's time to create some?
					</motion.p>
				</div>
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
			)}
		</>
	);
}

export default List;
