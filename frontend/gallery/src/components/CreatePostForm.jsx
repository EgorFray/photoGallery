import { usePosts } from "../context/PostsContext";
import { useNavigate } from "react-router-dom";
import styles from "./CreatePostForm.module.css";
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

				<div className={styles.row}>
					<label htmlFor="image">Add picture</label>
					<input type="file" className={styles.formInput} name="image" id="image" />
				</div>

				<div className={styles.row}>
					<label htmlFor="description">Add description</label>
					<textarea
						id="description"
						className={styles.formInput}
						name="description"
						placeholder="Add description"
					/>
				</div>

				<div className={styles.btnWrapper}>
					<Button className={styles.btnForm} type="submit">
						Post
					</Button>
				</div>
			</form>
		</div>
	);
}

export default CreatePostForm;
