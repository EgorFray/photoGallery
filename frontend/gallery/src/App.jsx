import { useState, useCallback, useEffect } from "react";
import Masonry from "react-masonry-css";
import { PostsProvider, usePosts } from "./context/PostsContext";
import { debounce } from "lodash-es";

function NavBar() {
	return (
		<nav className="navbar">
			<div className="logo-box">
				<img src="../public/Logo.png" alt='' className="logo" />
			</div>
			<div className="text-box">
				<p className="memorize">Memorize</p>
			</div>
		</nav>
	);
}

function Header({ onOpen }) {
	return (
		<div className="header">
			<div className="header-box">
				<div className="header-box--left">
					<h1 className="header-text--left">So, what's new?</h1>
				</div>

				<div className="header-box--right">
					<h2 className="header-text--right">Only your memories</h2>
					<button className="post-button" onClick={onOpen}>
						Post
					</button>
				</div>
			</div>
		</div>
	);
}

function Search() {
	const { getSearchedPosts } = usePosts();
	const [query, setQuery] = useState("");

	const debouncedSearch = useCallback(debounce(getSearchedPosts, 800), []);

	useEffect(() => {
		debouncedSearch(query)
	}, [query]);

	return (
		<div className="search">
			<input
				type="text"
				className="search-input"
				placeholder="Search your memory"
				onChange={(e) => setQuery(e.target.value)}
			/>
		</div>
	);
}

function CreatePostForm({ onOpen }) {
	const { createPost } = usePosts();

	async function handleSubmit(e) {
		e.preventDefault();

		const formData = new FormData(e.target);
		await createPost(formData);
		onOpen();
	}

	return (
		<div className="popup-overlay">
			<div className="popup">
				<form className="popup-form" onSubmit={handleSubmit}>
					<h2 className="popup-heading">Add your memory</h2>

					<button className="close-popup" onClick={onOpen}>
						x
					</button>

					<label className="popup-image">Add picture</label>
					<input type="file" className="images-val" name="image" />

					<label htmlFor="description">Description</label>
					<textarea
						id="description"
						className="description"
						name="description"
						placeholder="Add description"
					/>

					<button type="submit" className="button-submit">
						Post
					</button>
				</form>
			</div>
		</div>
	);
}

function Main({ children }) {
	return <section>{children}</section>;
}

function List() {
	const { posts } = usePosts();

	const breakpointColumnsObj = {
		default: 4,
		1100: 3,
		700: 2,
		500: 1,
	};

	return posts ? (
		<Masonry
			breakpointCols={breakpointColumnsObj}
			className="my-masonry-grid"
			columnClassName="my-masonry-grid_column"
		>
			{[...posts].reverse().map((post) => (
				<li key={post.ID} className="post">
					<img
						className="post-img"
						src={`http://localhost:8080${post.image}`}
						alt="Something beeautiful here"
					/>
					{/* HERE WE'LL ADD DATE IN THE FUTURE */}
					<div className="post-box">
						<p className="post-desc">{post.description} </p>
					</div>
				</li>
			))}
		</Masonry>
	) : (
		<div className="no-posts-wrapper">
			<p className="no-posts">There are no posts 🥲</p>
		</div>
	);
}

function Footer() {
	return (
		<footer className="footer">
			<p className="footer-text">Live life. Be creative</p>
		</footer>
	);
}

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
