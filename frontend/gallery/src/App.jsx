import { useState } from "react";
import { PostsProvider } from "./context/PostsContext";
import NavBar from "./components/NavBar";
import Header from "./components/Header";
import Search from "./components/Search";
import CreatePostForm from "./components/CreatePostForm";
import Main from "./components/Main";
import List from "./components/List";
import Footer from "./components/Footer";

function App() {
	const [isOpen, setIsOpen] = useState(false);

	function toggleForm() {
		setIsOpen(!isOpen);
	}

	return (
		<div>
			<PostsProvider>
				<NavBar />
				<Header onOpen={toggleForm} />
				<Search />
				{isOpen && <CreatePostForm onOpen={toggleForm} />}
				<Main>
					<List />
				</Main>
				<Footer />
			</PostsProvider>
		</div>
	);
}

export default App;
