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
		navigate("/");
	}

	if (isLoading) return <Spinner />;

	return (
		<div className={styles.formContainer}>
			<form className={styles.createForm} onSubmit={handleSubmit}>
				<h2 className={styles.createFormHeading}>Add your memory</h2>

				<div className={styles.imageContainer}>
					<label className={styles.createFormImage}>Add picture</label>
					<input type="file" name="image" />
				</div>

				<div className={styles.descriptionContainer}>
					<label className={styles.createFormDescription}>Add description</label>
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
