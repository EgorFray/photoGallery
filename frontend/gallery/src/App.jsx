import { useEffect } from "react";
import { useState } from "react";
import Masonry from "react-masonry-css";

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
	const [query, setQuery] = useState("");
	return (
		<div className="search">
			<input
				type="text"
				className="search-input"
				placeholder="Search your memory"
				onChange={(e) => setQuery(e.target.value)}
			/>
			<button type="submit" className="search-button">
				X
			</button>
		</div>
	);
}

function CreatePostForm({ onOpen, setPosts }) {
	async function handleSubmit(e) {
		e.preventDefault();

		const formData = new FormData(e.target);
		await createPost(formData);
		onOpen();
	}

	async function createPost(newPost) {
		try {
			const res = await fetch("http://localhost:8080/posts", {
				method: "POST",
				body: newPost,
			});
			const data = await res.json();
			setPosts((posts) => [...posts, data]);
		} catch {
			alert("There was an error loading data");
		}
	}

	return (
		<div className="popup-overlay">
			<div className="popup">
				<form className="popup-form" onSubmit={handleSubmit}>
					<h2 className="popup-heading">Add your memory</h2>

					<button class="close-popup" onClick={onOpen}>
						x
					</button>

					<label className="popup-image">Add picture</label>
					<input type="file" className="images-val" name="image" />

					<label for="description">Description</label>
					<textarea
						id="description"
						className="description"
						name="description"
						placeholder="Add description"
					></textarea>

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

function List({ posts, setPosts, setError }) {
	useEffect(function () {
		async function fetchGetData() {
			try {
				setError("");
				const res = await fetch("http://localhost:8080/posts");
				console.log(res);
				if (!res.ok) throw new Error("Something went wrong whil fetching data");
				const data = await res.json();
				console.log(data);
				setPosts(data);
			} catch (err) {
				if (err.name !== "AbortError") {
					setError(err.message);
				}
			}
			setError("");
		}
		fetchGetData();
	}, []);

	const breakpointColumnsObj = {
		default: 4,
		1100: 3,
		700: 2,
		500: 1,
	};

	return (
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
	);
}

function App() {
	const [isOpen, setIsOpen] = useState(false);

	const [posts, setPosts] = useState([]);
	const [error, setError] = useState("");

	function toggleForm() {
		setIsOpen(!isOpen);
	}

	return (
		<div>
			<Header onOpen={toggleForm} />
			<Search />
			{isOpen && <CreatePostForm onOpen={toggleForm} setPosts={setPosts} />}
			<Main>
				<List posts={posts} setPosts={setPosts} setError={setError} />
			</Main>
		</div>
	);
}

export default App;
