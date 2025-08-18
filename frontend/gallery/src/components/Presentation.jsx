import { motion } from "motion/react";

import Button from "./Button";
import styles from "./Presentation.module.css";

function Presentation() {
	const MotionButton = motion(Button);
	return (
		<div className={styles.container}>
			<div className={styles.textBox}>
				<motion.p
					className={styles.text}
					initial={{ clipPath: "inset(0 100% 0 0)" }}
					animate={{ clipPath: "inset(0 0% 0 0)" }}
					transition={{ duration: 1.2, ease: "easeInOut", delay: 2.5 }}
				>
					Create your own beautifull gallery
				</motion.p>
			</div>
			<div className={styles.imgBox}>
				<div className={styles.imgWrapper}>
					<motion.img
						src="/green.jpg"
						initial={{ opacity: 0, scale: 0.9 }}
						animate={{ opacity: 1, scale: 1 }}
						transition={{ duration: 0.8, ease: "easeOut", delay: 1 }}
					/>
				</div>
				<div className={styles.imgWrapper}>
					<motion.img
						src="/mist.jpg"
						initial={{ opacity: 0, scale: 0.9 }}
						animate={{ opacity: 1, scale: 1 }}
						transition={{ duration: 0.8, ease: "easeOut", delay: 2 }}
					/>
				</div>
				<div className={styles.imgWrapper}>
					<motion.img
						src="/tent.jpg"
						initial={{ opacity: 0, scale: 0.9 }}
						animate={{ opacity: 1, scale: 1 }}
						transition={{ duration: 0.8, ease: "easeOut", delay: 1.5 }}
					/>
				</div>
			</div>
			<div className={styles.btnBox}>
				<MotionButton
					initial={{ opacity: 0, y: 50 }}
					animate={{ opacity: 1, y: 0 }}
					transition={{ duration: 2, ease: "easeInOut" }}
				>
					Login
				</MotionButton>
			</div>
		</div>
	);
}

export default Presentation;
