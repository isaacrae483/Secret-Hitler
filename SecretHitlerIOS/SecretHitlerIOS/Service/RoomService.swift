//
// Created by Isaac Rae on 3/21/23.
//

import Foundation
import Combine

class RoomService {
    let client: APIClient

    init(_ client: APIClient) {
        self.client = client
    }

    func available() -> AnyPublisher<[Room], NetworkServiceError> {
        client.request(.available)
    }
}