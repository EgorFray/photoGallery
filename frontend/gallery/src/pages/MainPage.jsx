import { useState } from "react";
import CreatePostForm from "../components/CreatePostForm";
import Footer from "../components/Footer";
import Header from "../components/Header";
import List from "../components/List";
import Main from "../components/Main";
import NavBar from "../components/NavBar";
import Search from "../components/Search";

function MainPage() {
	const [isOpen, setIsOpen] = useState(false);

	function toggleForm() {
		setIsOpen(!isOpen);
	}

	return (
		<div>
			<NavBar />
			<Header onOpen={toggleForm} />
			{isOpen && <CreatePostForm onOpen={toggleForm} />}
			<Main>
				<List />
			</Main>
			<Footer />
		</div>
	);
}

export default MainPage;
