import Button from "./Button";
import styles from "./UpdateUserForm.module.css";

function UpdateUserForm() {
	return (
		<div className={styles.formContainer}>
			<p className={styles.updateHeading}>
				You don't have to update all field at once
			</p>
			<span className={styles.updateHeadingMotion}>Update only neccessary</span>
			<form className={styles.updateForm}>
				<div className={styles.row}>
					<label htmlFor="name">Name</label>
					<input
						type="text"
						id="name"
						name="name"
						// onChange={(e) => setName(e.target.value)}
						// value={name}
					/>
				</div>

				<div className={styles.row}>
					<label htmlFor="password">Password</label>
					<input
						type="password"
						id="password"
						name="password"
						// onChange={(e) => setPassword(e.target.value)}
						// value={password}
					/>
				</div>

				<div className={styles.row}>
					<label htmlFor="avatar" className={styles.addAvatar}>
						Add avatar
					</label>
					<input type="file" name="avatar" id="avatar" />
				</div>
				<Button className={styles.btnForm} type="submit">
					Update user
				</Button>
			</form>
		</div>
	);
}

export default UpdateUserForm;
