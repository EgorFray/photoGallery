import { useContext, useState } from "react";
import { createContext } from "react";

const PostDetailContext = createContext();

function PostDetailProvider({ children }) {
	const [post, setPost] = useState({});

	async function getPostById(id) {
		try {
			const res = await fetch(`http://localhost:8080/posts/${id}`);
			const data = await res.json();
			setPost(data);
		} catch {
			alert("There was an error loading post");
		}
	}
	return (
		<PostDetailContext.Provider value={{ post, getPostById }}>
			{children}
		</PostDetailContext.Provider>
	);
}

function usePostDetail() {
	const context = useContext(PostDetailContext);
	if (context === undefined)
		throw new Error("PostDetail context was used outside PostDetailProvider");
	return context;
}

export { PostDetailProvider, usePostDetail };
