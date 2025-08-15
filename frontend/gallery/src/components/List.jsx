import { usePosts } from "../context/PostsContext";
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

	if (isLoading) return <Spinner />;

	return posts ? (
		<Masonry
			breakpointCols={breakpointColumnsObj}
			className={styles.myMasonryGrid}
			columnClassName={styles.myMasonryGridColumn}
		>
			{[...posts].reverse().map((post) => (
				<Post post={post} />
			))}
		</Masonry>
	) : (
		<div className="noPostsWrapper">
			<p className="noPosts">There are no posts 🥲</p>
		</div>
	);
}

export default List;
