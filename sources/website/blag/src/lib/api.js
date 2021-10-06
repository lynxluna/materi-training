const API_URL = "http://127.0.0.1:8000"
const build_url = (path) => API_URL + path

const API = {
  create: async (title, content) => fetch(build_url('/articles'), {
    method: 'POST',
    credentials: 'include',
    headers: {
      'Content-Type':'application/json'
    },
    body: JSON.stringify({
      title: title,
      content: content,
    })
  }),

  single: async(id) => fetch(build_url('/articles/' + id), {
    method: 'GET',
    credentials: 'include',
  }),

  list_articles: async() => fetch (build_url('/articles'), {
    method: 'GET',
    credentials: 'include',
  })
}

export {API}
