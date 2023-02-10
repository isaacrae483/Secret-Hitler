//
// Created by Isaac Rae on 2/4/23.
//

import Foundation
import SwiftUI

struct RoundedButtonStyle: ButtonStyle {
    func makeBody(configuration: Self.Configuration) -> some View {
        configuration.label
                .padding()
                .frame(maxWidth: .infinity)
                .overlay(
                        RoundedRectangle(cornerRadius: 20)
                                .stroke(Color.black, lineWidth: 2)
                )
    }
}

extension ButtonStyle where Self == RoundedButtonStyle {
    static var rounded: Self {
        return .init()
    }
}