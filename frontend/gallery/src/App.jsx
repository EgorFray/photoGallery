import { PostsProvider } from "./context/PostsContext";
import PostDetail from "./pages/PostDetail";
import { Routes, Route } from "react-router-dom";
import MainPage from "./pages/MainPage";
import CreatePost from "./pages/CreatePost";
import PageNotFound from "./pages/PageNotFound";
import Homepage from "./pages/Homepage";
import Login from "./pages/Login";
import { AuthProvider } from "./context/FakeAuthContext";
import ProtectedRoute from "./pages/ProtectedRoute";

function App() {
	return (
		<AuthProvider>
			<PostsProvider>
				<Routes>
					<Route index element={<Homepage />} />
					<Route path="/login" element={<Login />} />
					<Route
						path="/app"
						element={
							<ProtectedRoute>
								<MainPage />
							</ProtectedRoute>
						}
					/>
					<Route
						path="/create"
						element={
							<ProtectedRoute>
								<CreatePost />
							</ProtectedRoute>
						}
					/>
					<Route
						path="/app/posts/:id"
						element={
							<ProtectedRoute>
								<PostDetail />
							</ProtectedRoute>
						}
					/>
					<Route path="*" element={<PageNotFound />} />
				</Routes>
			</PostsProvider>
		</AuthProvider>
	);
}

export default App;
