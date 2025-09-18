import Button from "./Button";
import styles from "./CreateUserForm.module.css";

function CreateUserForm() {
	return (
		<div className={styles.formContainer}>
			<form className={styles.createForm}>
				<h2 className={styles.createUserHeading} />

				<div className={styles.row}>
					<label htmlFor="email">Email</label>
					<input type="email" id="email" />
				</div>

				<div className={styles.row}>
					<label htmlFor="password">Password</label>
					<input type="password" id="password" />
				</div>

				<div className={styles.row}>
					<label htmlFor="avatar" className={styles.addAvatar}>
						Add avatar
					</label>
					<input type="file" name="avatar" id="avatar" />
				</div>
			</form>

			<Button>Create user</Button>
		</div>
	);
}

export default CreateUserForm;
