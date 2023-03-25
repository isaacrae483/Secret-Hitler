//
// Created by Isaac Rae on 3/18/23.
//

import Foundation
import Combine

enum Endpoint: String {
    case test = "test"
    case login = "users/login"
    case signup = "users/signup"

    case available = "rooms/available"
}

final class APIClient {
    let decoder = JSONDecoder()
    let encoder = JSONEncoder()

    init() {
        decoder.dateDecodingStrategy = .iso8601
    }

    //2023-03-02T23:14:55.99293-07:00

    let urlSession: URLSession = .shared
    let baseURLString: String = "http://localhost:8080/"

    func request<T: Decodable, K: Encodable>(_ endpoint: Endpoint, _ method: String = "GET", _ id: String? = nil, body: K = "") -> AnyPublisher<T, NetworkServiceError> {

        guard let url = URL(string: baseURLString + endpoint.rawValue) else {
            return Fail(error: NetworkServiceError.invalidURL).eraseToAnyPublisher()
        }

        print(readCookie(forURL: url))

        var request = URLRequest(url: url)
        request.setValue("application/json", forHTTPHeaderField: "Accept")
        request.httpMethod = method

        if method != "GET" {
            let jsonEncoder = JSONEncoder()
            guard let jsonResultData = try? jsonEncoder.encode(body) else {
                return Fail(error: NetworkServiceError.invalidURL).eraseToAnyPublisher()
            }
            request.httpBody = jsonResultData
        }

        return urlSession.dataTaskPublisher(for: request)
                .tryMap { (data, response) -> Data in
                    if let httpResponse = response as? HTTPURLResponse {
                        print(httpResponse.statusCode)
                        guard (200..<300) ~= httpResponse.statusCode else {
                            throw NetworkServiceError.invalidResponseCode(httpResponse.statusCode)
                        }
                    }
                    return data
                }
                .decode(type: T.self, decoder: decoder)
                .mapError { error -> NetworkServiceError in
                    if let decodingError = error as? DecodingError {
                        return NetworkServiceError.decodingError((decodingError as NSError).debugDescription)
                    }
                    return NetworkServiceError.genericError(error.localizedDescription)
                }
                .eraseToAnyPublisher()
    }

}

func readCookie(forURL url: URL) -> [HTTPCookie] {
    let cookieStorage = HTTPCookieStorage.shared
    let cookies = cookieStorage.cookies(for: url) ?? []
    return cookies
}

func deleteCookies(forURL url: URL) {
    let cookieStorage = HTTPCookieStorage.shared

    for cookie in readCookie(forURL: url) {
        cookieStorage.deleteCookie(cookie)
    }
}

func storeCookies(_ cookies: [HTTPCookie], forURL url: URL) {
    let cookieStorage = HTTPCookieStorage.shared
    cookieStorage.setCookies(cookies,
            for: url,
            mainDocumentURL: nil)
}

enum NetworkServiceError: Error {
    case invalidURL
    case decodingError(String)
    case genericError(String)
    case invalidResponseCode(Int)

    var errorMessageString: String {
        switch self {
        case .invalidURL:
            return "Invalid URL encountered. Can't proceed with the request"
        case .decodingError:
            return "Encountered an error while decoding incoming server response. The data couldn’t be read because it isn’t in the correct format."
        case .genericError(let message):
            return message
        case .invalidResponseCode(let responseCode):
            return "Invalid response code encountered from the server. Expected 200, received \(responseCode)"
        }
    }
}