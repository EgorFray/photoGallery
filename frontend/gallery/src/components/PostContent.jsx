import { useEffect } from "react";
import { useParams } from "react-router-dom";
import { usePostDetail } from "../context/PostDetailContext";
import styles from "./PostContent.module.css";

function PostContent() {
	const { id } = useParams();
	const { post, getPostById } = usePostDetail();

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

	console.log(JSON.stringify(post));

	return (
		<div className={styles.detailLayout}>
			<img
				src={`http://localhost:8080${post.image}`}
				className={styles.detailImage}
			/>
			<div className={styles.detailBox}>
				<p className={styles.detailDescription}>{post.description}</p>
				<div className={styles.boxBottom}>
					<p className={styles.detailDate}>{formatDate(post.created_at)}</p>
					<button className={styles.deleteButton}>Delete</button>
				</div>
			</div>
		</div>
	);
}

export default PostContent;
