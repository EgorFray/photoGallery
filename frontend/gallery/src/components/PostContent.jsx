import { useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { usePosts } from "../context/PostsContext";
import styles from "./PostContent.module.css";
import Button from "./Button";
import Spinner from "./Spinner";

function PostContent() {
	const { id } = useParams();
	const { post, getPostById, isLoading, deletePost } = usePosts();
	const navigate = useNavigate();

	const formatDate = (date) =>
		new Intl.DateTimeFormat("en", {
			day: "numeric",
			month: "long",
			year: "numeric",
			weekday: "long",
		}).format(new Date(date));

	useEffect(
		function () {
			getPostById(id);
		},
		[id]
	);

	function handleClick() {
		deletePost(id);
		navigate("/");
	}

	if (isLoading) return <Spinner />;

	return (
		<div className={styles.detailLayout}>
			<img
				src={`http://localhost:8080${post.image}`}
				className={styles.detailImage}
			/>
			<div className={styles.detailBox}>
				<p className={styles.detailDescription}>{post.description}</p>
				<div className={styles.boxBottom}>
					<p className={styles.detailDate}>{formatDate(post.created_at || null)}</p>
					<Button handleClick={handleClick}>Delete</Button>
				</div>
			</div>
		</div>
	);
}

export default PostContent;
