//
// Created by Isaac Rae on 2/4/23.
//

import Foundation
import SwiftUI

extension JoinView {
    @MainActor class ViewModel: ObservableObject {
        @Published var showSheet: Bool = false

        @Published var code: String = ""

        init() {}

        func joinRoom() {
            print("joining")
            showSheet.toggle()
            print(showSheet)
        }

        func createRoom() {
            print("creating")
        }

        func browseRoom() {
            print("browsing")
        }
    }
}