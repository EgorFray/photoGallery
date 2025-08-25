import { usePosts } from "../context/PostsContext";
import styles from "./CreatePostForm.module.css";
import { useNavigate } from "react-router-dom";
import Spinner from "./Spinner";
import Button from "./Button";

function CreatePostForm() {
	const { createPost, isLoading } = usePosts();
	const navigate = useNavigate();

	async function handleSubmit(e) {
		e.preventDefault();

		const formData = new FormData(e.target);
		await createPost(formData);
		navigate("/app");
	}

	if (isLoading) return <Spinner />;

	return (
		<div className={styles.formContainer}>
			<form className={styles.createForm} onSubmit={handleSubmit}>
				<h2 className={styles.createFormHeading}>Add your memory</h2>

				<div className={styles.imageContainer}>
					<label htmlFor="image" className={styles.createFormImage}>
						Add picture
					</label>
					<input type="file" name="image" id="image" />
				</div>

				<div className={styles.descriptionContainer}>
					<label htmlFor="description" className={styles.createFormDescription}>
						Add description
					</label>
					<textarea
						id="description"
						name="description"
						placeholder="Add description"
					/>
				</div>

				<Button type="submit">Post </Button>
			</form>
		</div>
	);
}

export default CreatePostForm;
