import { usePosts } from "../context/PostsContext";
import styles from "./CreatePostForm.module.css";

function CreatePostForm({ onOpen }) {
	const { createPost } = usePosts();

	async function handleSubmit(e) {
		e.preventDefault();

		const formData = new FormData(e.target);
		await createPost(formData);
		onOpen();
	}

	return (
		<div className={styles.popupOverlay}>
			<div className={styles.popup}>
				<form className={styles.popupForm} onSubmit={handleSubmit}>
					<h2 className={styles.popupHeading}>Add your memory</h2>

					<button className={styles.closePopup} onClick={onOpen}>
						x
					</button>

					<label className={styles.popupImage}>Add picture</label>
					<input type="file" className={styles.imagesVal} name="image" />

					<label htmlFor={styles.description}>Description</label>
					<textarea
						id="description"
						className={styles.description}
						name="description"
						placeholder="Add description"
					/>

					<button type="submit" className={styles.buttonSubmit}>
						Post
					</button>
				</form>
			</div>
		</div>
	);
}

export default CreatePostForm;
