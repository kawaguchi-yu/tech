export const sessionInformation = process.env.NODE_ENV === 'development' ? {
  backendHost: 'http://localhost:8080'
} :
  {
  backendHost: 'http://backend.quiztecher.com'
}