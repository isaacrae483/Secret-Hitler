//
// Created by Isaac Rae on 3/21/23.
//

import Foundation
import SwiftUI

final class Service {
    @Published var loggedIn = false
    @Published var roomCode = ""
    let client: APIClient = APIClient()
    let rooms: RoomService
    let auth: AuthenticationService

    init() {
        rooms = RoomService(client)
        auth = AuthenticationService(client)
    }
}